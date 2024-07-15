#include "triangle.h"

static bool is_triangle(triangle_t triangle)
{
    const double a = triangle.a;
    const double b = triangle.b;
    const double c = triangle.c;

    return (a > 0 && b > 0 && c > 0)
        && ((a + b) >= c)
        && ((b + c) >= a)
        && ((a + c) >= b);
}

bool is_equilateral(triangle_t triangle)
{
    const double a = triangle.a;
    const double b = triangle.b;
    const double c = triangle.c;

    return is_triangle(triangle) && a == b && a == c;
}

bool is_isosceles(triangle_t triangle)
{
    const double a = triangle.a;
    const double b = triangle.b;
    const double c = triangle.c;

    return is_triangle(triangle) && (a == b || a == c || b == c);
}

bool is_scalene(triangle_t triangle)
{
    return is_triangle(triangle)
        && !is_equilateral(triangle)
        && !is_isosceles(triangle);
}
