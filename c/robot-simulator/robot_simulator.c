#include "robot_simulator.h"

robot_status_t robot_create(robot_direction_t direction, int x, int y)
{
    robot_position_t position = { x, y };
    robot_status_t status = { direction, position };
    return status;
}

static void move(robot_status_t *robot)
{
    switch (robot->direction) {
        case DIRECTION_NORTH:
            robot->position.y++;
            break;
        case DIRECTION_EAST:
            robot->position.x++;
            break;
        case DIRECTION_SOUTH:
            robot->position.y--;
            break;
        case DIRECTION_WEST:
            robot->position.x--;
            break;
        default:
            // do nothing
            break;
    }
}

static void rotate_left(robot_status_t *robot)
{
    switch (robot->direction) {
        case DIRECTION_NORTH:
            robot->direction = DIRECTION_WEST;
            break;
        case DIRECTION_EAST:
            robot->direction = DIRECTION_NORTH;
            break;
        case DIRECTION_SOUTH:
            robot->direction = DIRECTION_EAST;
            break;
        case DIRECTION_WEST:
            robot->direction = DIRECTION_SOUTH;
            break;
        default:
            // do nothing
            break;
    }
}

static void rotate_right(robot_status_t *robot)
{
    switch (robot->direction) {
        case DIRECTION_NORTH:
            robot->direction = DIRECTION_EAST;
            break;
        case DIRECTION_EAST:
            robot->direction = DIRECTION_SOUTH;
            break;
        case DIRECTION_SOUTH:
            robot->direction = DIRECTION_WEST;
            break;
        case DIRECTION_WEST:
            robot->direction = DIRECTION_NORTH;
            break;
        default:
            // do nothing
            break;
    }
}

void robot_move(robot_status_t *robot, const char *commands)
{
    for (char c; (c = *commands); commands++) {
        switch (c) {
            case 'A':
                move(robot);
                break;
            case 'L':
                rotate_left(robot);
                break;
            case 'R':
                rotate_right(robot);
                break;
            default:
                // do nothing
                break;
        }
    }
}
