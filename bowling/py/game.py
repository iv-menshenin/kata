class Bowling:

    def __init__(self):
        self.is_spare = False
        self.is_strike = 0
        self.roll_num = 0
        self.curr_set = 0
        self.sum_scores = 0

    def score(self):
        return self.sum_scores + self.curr_set

    def roll(self, score):
        self._roll_calculate(score)
        self._bonus_calculate(score)
        if self._is_set_over():
            self._init_new_set()

    def _roll_calculate(self, score):
        self.roll_num = self.roll_num + 1
        self.curr_set = self.curr_set + score

    def _bonus_calculate(self, score):
        if self.is_strike > 0:
            if self.is_strike > 2:
                self.sum_scores = self.sum_scores + score
                self.is_strike = self.is_strike - 1
            self.sum_scores = self.sum_scores + score
            self.is_strike = self.is_strike - 1
        if self.is_spare:
            self.sum_scores = self.sum_scores + score
        self.is_spare = False

    def _is_set_over(self):
        return self.roll_num == 2 or self.curr_set == 10

    def _init_new_set(self):
        if self._is_strike():
            self.is_strike = self.is_strike + 2
        self.is_spare = self._is_spare()
        self.sum_scores = self.sum_scores + self.curr_set
        self.curr_set = 0
        self.roll_num = 0

    def _is_strike(self):
        return self.curr_set == 10 and self.roll_num == 1

    def _is_spare(self):
        return self.curr_set == 10 and self.roll_num == 2
