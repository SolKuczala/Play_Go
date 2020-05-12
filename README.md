# Tic Tac Toe Game
Implements package, cli and api

> cd tic-tac-toe  
> docker build .   

(optional tag)

Two options, changing env var "MODE", to run either the api:
> docker run -e MODE=api -p {port you want}:9090 -it {id or docker tag}

or the cli:
> docker run -e MODE=cli -it {id or docker tag}

Api use:

with curl, hit your local machine to the port binded with the next paths:

GET {ip:port}/create-board/{number}   
to create a new board with that size (is set between 2 and 9)

PUT{ip:port}/send-play/{X or O}/{row number}/{column number}  
to send the player and the coordinates where to place that play

GET{ip:port}/status
to see status of the game (board, last player that played)

Once ended, must get another board in order to start again.
