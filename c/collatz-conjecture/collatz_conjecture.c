#include "collatz_conjecture.h"

/*
 * steps: returns the number of steps Collatz's conjecture takes to get from
 * 	the value provided to the number 1, or returns -1 if the value provided
 * 	is less than 1.
 *
 * See the Collatz conjecture wikipedia page for more information:
 * 	https://en.wikipedia.org/wiki/Collatz_conjecture
 */
int steps(int value)
{
	if (value < 1) {
		return ERROR_VALUE;
	}

	unsigned int nsteps;
	for (nsteps = 0; value != 1; nsteps++) {
		if (value%2 == 0) {
			value /= 2;
		} else {
			value = (3 * value) + 1;
		}
	}

	return nsteps;
}
