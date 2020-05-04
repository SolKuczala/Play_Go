# Tic Tac Toe Game
Implements package, cli and api

> cd tic-tac-toe  
> docker build .   

(optional tag)

Two options, changing env var "MODE", to run either the api:
> docker run -e MODE=api -p {port you want}:9090 -it {id or docker tag}

or the cli:
>docker run -e MODE=cli -it {id or docker tag}

Api use:

with curl, hit your local machine to the port binded to the next paths:

/create-board/{number}   
to create a new board with that size (is set between 2 and 9)

/send-play/{X or O}/{row nomber}/{column number}  
to send the player and the coordinates where to place that play

<<<<<<< HEAD
Once ended, must get another board in order to start again.
=======
Once ended must send get another board in order to start again.
>>>>>>> 8114fc22c390fd8fd7674c094a29fb59f51d9d7e
