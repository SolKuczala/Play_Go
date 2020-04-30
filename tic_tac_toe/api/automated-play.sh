#automated play
curl -GET http://127.0.0.1:9090/create-board/3 | jq

curl -XPUT http://127.0.0.1:9090/send-play/O/0/0 | jq 
curl -XPUT http://127.0.0.1:9090/send-play/X/1/1 | jq
curl -XPUT http://127.0.0.1:9090/send-play/O/0/1 | jq
curl -XPUT http://127.0.0.1:9090/send-play/X/1/2 | jq
curl -XPUT http://127.0.0.1:9090/send-play/O/0/2 | jq
