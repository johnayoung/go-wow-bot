// #include "stdafx.h"
#include <iostream>
#include <windows.h>

struct Player
{
    int level;
    int health;

    // new yell function
    void yell()
    {
        std::cout << "Raaghgghh!!!!" << std::endl;
    }
};

struct Game
{
    Player *player;
};

Game *pGame;
int main()
{
    pGame = new Game;
    pGame->player = new Player;
    pGame->player->level = 1;
    pGame->player->health = 100;

    INPUT_RECORD event;
    HANDLE hStdIn = GetStdHandle(STD_INPUT_HANDLE);
    DWORD count;

    std::cout << "Welcome to Bloog's Quest!" << std::endl;
    std::cout << "Player is level 1 and has 100 health." << std::endl
              << std::endl;

    while (true)
    {
        // handle user input
        if (WaitForSingleObject(hStdIn, 0) == WAIT_OBJECT_0)
        {
            ReadConsoleInput(hStdIn, &event, 1, &count);

            if (event.EventType == KEY_EVENT && !event.Event.KeyEvent.bKeyDown)
            {
                switch (event.Event.KeyEvent.wVirtualKeyCode)
                {
                case VK_SPACE:
                    // call the new yell function when the spacebar is pressed
                    pGame->player->yell();
                    pGame->player->health--;
                    std::cout << "After taking 1 damage, Player's remaining health is: "
                              << pGame->player->health << std::endl;
                    break;
                case VK_ESCAPE:
                    return 0;
                }
            }
        }

        // update simulation

        // render graphics
    }
}