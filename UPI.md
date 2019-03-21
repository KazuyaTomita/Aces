## Overview

UPI is a protocol which separates interface parts and actual thinking parts.
And the documentation of UPI is [here](https://cdn.shopify.com/s/files/1/0769/9693/files/UPI19-documentation.pdf?9664768544248083387).
However, the protocol is designed for people to give a detailed situation which they want to dive deep. People gives a lot of information of the situation and then calculation is executed.
Even thought it is useful, we want to build a fully-thinking engine which needs only a present situation and history.
Therefore we need to change the protocol like Universal Shogi Interface(USI) or Universal Chess Interface(UCI).
Here are some commands.

### Commands from GUI to an engine

upi
    
    GUI sends this command when GUI starts an engine. 
    The engine should send `upiok`.
    
isready

    Useful for checking. The engine should send `is_readyok`
    
upinewgame

    A sign of starting a new game.
    
state

    GUI sends a present situation, including a board, action history, bet history and something like that.
    The expression of a state is inspired by the ACPC protocol. 
    Note that the modified UPI doesn't need to follow the ACPC, but the expression is good.
    So, we use this expression. The summary of the expression is like this.
    <position>:<game id>:<action history>:<cards>|<boardCards>:<action>
    ex)
    GUI ->   state 0:0:r20r50c/rc/:TdAs|/2c8c3h/9c 
    Engine -> action 0:0:r20r50c/rc/:TdAs|/2c8c3h/9c:c 
 
go

    A sign of starting thinking. An engine starts thinking.
    
exit

    kill an engine.
    
gameover

    The game is over. The result sends to an engine.
    
stop

    If an engine receives stop, the engine should send a best action at the time.
    
### Commands from an engine to GUI

upiok
    
    A command which is correponding to `upi`.
    
is_readyok

    A command which is correponding to `isready`. Initialization should be done until sending `is_readyok`.
    
action

    send engine's action. This command send of course an action and the state.
    
    ex) do call(check) at turn
    action 0:0:r20r50c/rc/:TdAs|/2c8c3h/9c:c 
    

    
    



