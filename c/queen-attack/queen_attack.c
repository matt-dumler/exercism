#include "queen_attack.h"

// is_valid_position: the position is within the boundaries of the board
static bool is_valid_position(position_t pos)
{
    return pos.row < 8 && pos.column < 8;
}

/*
 * can_attack: checks if two queens could attack each other based on their
 *     board position.
 */
attack_status_t can_attack(position_t queen_1, position_t queen_2)
{
    if (!is_valid_position(queen_1) || !is_valid_position(queen_2)) {
        return INVALID_POSITION;
    } else if (queen_1.row == queen_2.row && queen_1.column == queen_2.column) {
        return INVALID_POSITION;
    } else if (queen_1.row == queen_2.row || queen_1.column == queen_2.column) {
        return CAN_ATTACK;
    } else if (abs(queen_1.row - queen_2.row) == abs(queen_1.column - queen_2.column)) {
        return CAN_ATTACK; // attack across diagonals
    }

    return CAN_NOT_ATTACK;
}
