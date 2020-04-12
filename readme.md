Цель:
1) Разобраться с основными элементами блокчейна
 * Блок
 * Транзакции и подписи
 * Майнинг
 * Нетворкинг
 * Работа с нодами
 * Мехаизм синхронизации нод
 * Апи клиента
   
2) Что включает:
 - Консенсус
 - Майнинг
 - Пул транзакций
 - Работа с сетью
 - Выполнение блоков
 - Синк

 
3) Заготовка
Что входит в заготовку
 * Набор интерфейсов
 * Базовый тест

4) Задание:
     - Начальное состояние блокчейна(Genesis)
        - В начальное состояние входят
            - Список валидаторов
            - Список адресов на которые будут начислены деньги в первом блоке
            - Genesis преобразуется в первый блок каждой ноды отдельно при старте ноды

     - Есть 2 вида пользователей:
         - майнеры, которые создают блоки. У них есть пул транзакций. Майнеры в то же время являются и обычными пользователями.
         - обычные пользователи, которые умеют отправлять транзакции от своего имени и которые должны пересылать всем своим пирам пришедшие транзакции других пользователей.

     - Хранилище неперсистентное, но задание со * - персистентное.

     - Для каждого пира должен быть сгенерирован приватный ключ
        - Им валидаторы подписывают блок
        - Хеш публичного ключа от этого приватного используется в качестве имени пира и адреса кошелька(преобразованый в hex)
        
     - Реализовать механизм соединения пиров через каналы
        - При подключении происходит "handshake"
            - Оба пира посылают номер последнего блока(lastBlockNum)
                - и td и хэш последнего блок для задания PoW
        - Если lastBlockNum пира больше локального lastBlockNum, то скачиваем недостающие блоки(см Алгоритм синхронизации)
        - Если нода валидатор, то она создает новый блок - после его создания она рассылает его всем подключенным пирам
        - При получении нового блока и после успешного выполнения его на ноде, нода должна отправить его всем, кроме того, кто его прислал
        - Если нода получила блок, о котором она уже знает - нода ничего не должна делать  
        
     - Алгоритм консенсуса(Proof of Autority)
        - Валидаторы, список которых определен в Genesis блоке, создают блоки
        - Каждый блок создается и подписывается одним валидатором
        - Очередность валидаторов определяется в Genesis блоке
        - Если отсутствует текущая валидирующая нода, то консенсус будет ждать ее появления
        - Тот, кто произвел блок получает награду в 1000 монет
        - При формировании нового блока, в него добаляются транзакции из пула транзакций
            - Не более 10 штук на блок
            - При выполнении транзакции
                - Отнимается баланс у From
                - Прибавляетс баланс To
                - Fee Добавляется к балансу валидатора
                - Если транзакция ошибочна(баланс, подпись) - не добавляет ее в блок
                - У отправителя должно быть достаточно баланса для полей Amont и Fee
                - ?? Что делать с невалидными транзакциями
        - После выполнения транзакций считается StateHash, который записывается в блок
        - Заполняются остальные поля блока
        - Считается хеш блока
        - Валидатор подписывает блок и отправляет его всем подключенным пирам
        - Каждый пир, получивший блок - проверяет
            - Подпись, что именно этот валидатор должен создавать блок
                - Что подписывал правильный валидатор в свою очередь
            - Что хеш блока корректный
            - Что хеш родителя корректный
            - Что все транзакции выполняются корректно
            - Что подпись транзакции корректна и валидна

    - Пул транзакций
        - Хранит список отправленных клиентом транзакций
        - Отправляет отправляет новые и новые пришедшие транзакции своим пирам(кроме тех, кто прислал транзакцию на ноду)
        - Отсюда берут валидаторы транзакции для добавления в блок
        - После выполнения блока из пула транзакции нужно удалить выполненные транзакции
        
    - Алгоритм синхронизации
        - Только что подключенный пир узнает номер самого последнего блока во время handshake 
        - Последовательно скачивает и выполняет блоки
        - Если приходят блок, у которого номер блока больше чем lastBlockNum+1 - сохраняем его в отдельное место
        - Когда скачали все блоки до сохраненных - выполняем их
        
        
Индивидуальные задания:
 - PoW https://github.com/ethereum/wiki/wiki/Mining https://m.habr.com/ru/post/320178/
    - Надо реалзовать простой PoW, где сложность - это количество лидирующих нулей хэша блока (difficulty). Такой хеш мы находим перебирая разные nonce блока. 
    - Должен быть реализован механизм переключения на новую canonical chain, canonical chain считается та, у которой total difficulty (TD) больше.
    todo - добавить алгоритм подбора difficulty
 
 - Легкий клиент https://link.medium.com/MgGCjPM4n5
    - Сделать новый вид нод - легкий клиент. Когда легкий клиент подключается к обычному - ему пересылается результат выполнения каждого блока(только изменение состояния после выполнения транзакций), подписаный приватным ключем блока и StateHash, который должен получиться.
    - Добавить новый запрос в апи - получить по номеру блока - StateHash и хеш блока.
    - Проверять каждые n блоков, что на n'й блок хеш блока и StateHash одинаковые ???  todo перефразировать
 
 - Хардфорк https://m.habr.com/ru/post/320178/
    - Реализовать механизм хардфорков, это способ, позволяющий менять протокол работы блокчейна. Реализуется через особый конфиг, где есть настройки, какой хардфорк (его имя), когда (номер блока) должен включаться. 
    - И рализовать один хардфорк с переходом консенсуса с подписи одного валидатора для блока на подпись 2х валидаторов(валидатор n -> валидатор n и n+1).
    - Self destoy валидатора. Каждый валидатор может отправить специальную транзакцию удаления себя из списка валидаторов. После ее обработки каждая нода должна удалить эту ноду из списка валидаторов.
 
 - Мультисиг транзакции https://m.habr.com/ru/company/mixbytes/blog/412675/
    - нужны "отложенные" транзакции. У них указывается кто платит, кому, сколько, не позже какого блока, и кто еще должен подтвердить этот перевод. 
    - Потом мы ждем, когда придут подтверждающие транзакции из списка. 
    - Как только соберутся все подписи, то начальная multisig транзакция выполняется.
    - Шаги
        - Блокировка средств отправителя
        - Еще транзакции с подписями
        - Когда набирается нужное число подписей - перевести средства
        - Если не набралось нужное количество подписей в течении n блоков - вернуть деньги отправителю(отдельная транзакция возврата)
    
 - Commit-reveal 
    - Надо реклизовать схему commit-reveal генерации случайного числа. 
    - Каждый участник должен прислать транзакцию с зашифрованным симметричным ключом случайным числом. 
    - Собираем 10 блоков, после чего ждем от участников объявления секретов. 
    - По истечению еще 10 блоков заканчиваем прием секретов. 
    - Каждый пользователь теперь может расшифровать все зашифрованные сообщения, к которым прислали секреты, конкатенировать их в том порядке, как приходили транзакции с commit, и захэшировать sha256, что и будет случайным числом.
    - Ссылки
        - https://m.habr.com/ru/post/348838/
        - https://m.habr.com/ru/post/448330/
        - https://m.habr.com/ru/post/452340/
