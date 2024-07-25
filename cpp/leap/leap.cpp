#include "leap.h"

#include <chrono>

namespace leap {

    bool is_leap_year(int year) {
        return std::chrono::year{year}.is_leap();
    }

}  // namespace leap
