#include "bob.h"

#include <ctype.h>
#include <stdbool.h>
#include <string.h>

static bool is_silence(char *greeting)
{
    for (char *p = greeting; *p != '\0'; p++) {
        if (!isspace(*p)) {
            return false;
        }
    }

    return true;
}

static bool is_question(char *greeting)
{
    char *p = greeting + (strlen(greeting) - 1);

    while (isspace(*p))
        p--;

    return *p == '?';
}

static bool is_shouted(char *greeting)
{
    bool is_shouted = false;
    for (char *p = greeting; *p != '\0'; p++) {
        if (isalpha(*p)) {
            if (isupper(*p)) {
                is_shouted = true;
            } else {
                return false;
            }
        }
    }

    return is_shouted;
}

char *hey_bob(char *greeting)
{
    if (is_silence(greeting)) {
        return "Fine. Be that way!";
    } else if (is_question(greeting) && is_shouted(greeting)) {
        return "Calm down, I know what I'm doing!";
    } else if (is_question(greeting)) {
        return "Sure.";
    } else if (is_shouted(greeting)) {
        return "Whoa, chill out!";
    } else {
        return "Whatever.";
    }
}
