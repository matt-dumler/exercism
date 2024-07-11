#include "high_scores.h"

#include <stdlib.h>
#include <string.h>

// latest: returns the last score in the list.
int32_t latest(const int32_t *scores, size_t scores_len)
{
    return scores[scores_len - 1];
}

// personal_best: returns the highest score in the list.
int32_t personal_best(const int32_t *scores, size_t scores_len)
{
    int32_t high_score = 0;
    for (size_t i = 0; i < scores_len; i++) {
        if (scores[i] > high_score) {
            high_score = scores[i];
        }
    }

    return high_score;
}

static int descending(const void *p, const void *q)
{
    return *(int *)q - *(int *)p;
}

/*
 * personal_top_three: copies up to the three highest scores to the user
 *     provided output pointer, and returns the number of scores copied to
 *     output.
 */
size_t personal_top_three(const int32_t *scores, size_t scores_len,
        int32_t *output)
{
    // Create a buffer and calculate needed valueas ahead of time.
    int32_t buffer[scores_len];
    size_t score_size = sizeof(int32_t),
           nscores = sizeof(scores) / score_size;

    // Create a copy of and sort the scores.
    memcpy(buffer, scores, scores_len * score_size);
    qsort((void *)buffer, nscores, score_size, descending);

    size_t count = nscores > 3 ? 3 : nscores;
    memcpy(output, buffer, count * score_size);
    return count;
}
