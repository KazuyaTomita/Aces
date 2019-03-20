## Overview

UPI is a protocol which separates interface parts and actual thinking parts.
And the documentation of UPI is [here](https://cdn.shopify.com/s/files/1/0769/9693/files/UPI19-documentation.pdf?9664768544248083387).
However, the protocol is designed for people to give a detailed situation which they want to dive deep. People gives a lot of information of the situation and then calculation is executed.
Even thought it is useful, we want to build fully-thinking engine which needs only a present situation and history.
Therefore we need to change the protocol like Universal Shogi Interface(USI) or Universal Chess Interface(UCI).
Here are some commands.

upi
    
    GUI sends this command when GUI starts an engine. 
    The engine should send `upi ok!`.
    
isready

    Useful for checking. The engine should send `is_ready ok!`
    
upinewgame

    A sign of starting a new game.
    
state

    GUI sends a present situation, including a board, action history, bet history and something like that.
 
go

    A sign of starting thinking. An engine starts thinking.
    
exit

    kill an engine.
    
gameover

    The game is over. The result sends to an engine.
    
stop

    If an engine receives stop, the engine should send a best action at the time.
    
    



