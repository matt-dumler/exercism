#include "difference_of_squares.h"

/*
 * sum_of_squares: returns 1 * 1 + 2 * 2 + 3 * 3 + ... + number * number
 *
 * For information on the formula used see:
 * 	https://www.geeksforgeeks.org/how-to-calculate-the-sum-of-squares/
 */
unsigned int sum_of_squares(unsigned int number)
{
	return (number * (number + 1) * ((number * 2) + 1)) / 6;
}

/*
 * square_of_sum: returns the square of 1 + 2 + 3 + ... + number
 *
 * For information on the formula used see:
 * 	https://www.geeksforgeeks.org/sum-of-arithmetic-sequence-formula/
 */
unsigned int square_of_sum(unsigned int number)
{
	return pow(((number * (number + 1)) / 2), 2);
}

unsigned int difference_of_squares(unsigned int number)
{
	return square_of_sum(number) - sum_of_squares(number);
}

