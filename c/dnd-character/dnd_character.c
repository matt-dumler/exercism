#include "dnd_character.h"

#include <math.h>
#include <stdlib.h>
#include <time.h>

/*
 * ability: generate a number between 3 and 18, which is equivalent to the sum
 *     of four dice rolls dropping the lowest result.
 */
int ability(void)
{
    srand(time(NULL));
    return (rand() % 18) + 3;
}

// modifier: returns an ability score minus ten and divided by two.
int modifier(int score)
{
    return floor(((float)score - 10.0) / 2.0);
}

dnd_character_t make_dnd_character(void)
{
    dnd_character_t character = {
        ability(), // strength
        ability(), // dexterity
        ability(), // constitution
        ability(), // intelligence
        ability(), // wisdom
        ability(), // charisma
        10         // base hitpoints
    };

    character.hitpoints += modifier(character.constitution);

    return character;
}
