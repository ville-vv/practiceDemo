let VIL = (function () {
    let VIL = {
    };



    function DefaultWebSocket(host, call) {
        let _host = host;
        let _isOpen = false;
        let _bufQueue = [];
        let _bufCap = 100;
        let _call = null;
        if("undefined" !== typeof call && call !== null){
            _call = call
        }else{
            _call = {
                onConnect:function (e) {
                    console.log("connect success ", e);
                },
                onDisconnect:function (e) {
                    console.log("disconnect ", e);
                },
                onMsg:function (data) {
                    //console.log("receive message ", data)
                }
            }
        }


        let _socket = new WebSocket(_host);
        _socket.binaryType = "arraybuffer";
        /**
         * 设置发送消息缓存队列的容量
         * @param {number} cap
         * @constructor
         */
        this.setBufferCap = function(cap){
            if("number" !== typeof cap ){
                console.error("parameter type is not number ");
                return ;
            }
            if(cap < 0){
                console.error("parameter value can not less then 0");
                return ;
            }
            _bufCap = cap;
        };

        /**
         * 发送消息
         * @param {string | ArrayBuffer } data
         * @constructor
         */
        this.send = function(data){
            if(_isOpen && _socket){
                _socket.send("");
            }else{
                if (_bufQueue < _bufCap){
                    _bufQueue.push(data);
                }
            }
        };

        this.close = function(){
            _socket.close(1000, "normal");
        };

        _socket.onopen = function(even){
            _isOpen = true;
            _call.onConnect(even);
            while (_bufQueue > 0){
                _socket.send(_bufQueue.shift());
            }
        };

        _socket.onmessage = function(e){
            let data = e.data;
            _call.onMsg(data);
        };

        /**
         * 收到关闭连接
         * @param even
         */
        _socket.onclose = function(even){
            _isOpen = false;
            _call.onDisconnect({host:_host, event:even});
        };

        /**
         * 收到错误
         * @param err
         */
        _socket.onerror = function(err){
            _isOpen = false;
            _call.onDisconnect({host:_host, event:err});
        };
    }

    try{
        VIL.EngineSocket = DefaultWebSocket ;
    }catch (e) {
        console.error("VILEngine error ", e);
    }

    return VIL;
})();
