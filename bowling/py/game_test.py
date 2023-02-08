import unittest

from game import Bowling


class TestFactor(unittest.TestCase):

    def setUp(self) -> None:
        self.bowling = Bowling()

    def test_play_zero(self):
        self.assertEqual(0, self.bowling.score())

    def test_roll_one(self):
        self.bowling.roll(1)
        self.assertEqual(1, self.bowling.score())

    def test_roll_twice(self):
        self.bowling.roll(1)
        self.bowling.roll(2)
        self.assertEqual(3, self.bowling.score())

    def test_spare(self):
        self.bowling.roll(5)
        self.bowling.roll(5)
        self.bowling.roll(3)
        self.assertEqual(16, self.bowling.score())

    def test_double_spare(self):
        self.bowling.roll(4)
        self.bowling.roll(6)
        self.bowling.roll(2)
        self.bowling.roll(8)
        self.bowling.roll(1)
        self.assertEqual(24, self.bowling.score())

    def test_strike(self):
        self.bowling.roll(10)
        self.bowling.roll(8)
        self.bowling.roll(1)
        self.assertEqual(28, self.bowling.score())

    def test_double_strike(self):
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(2)
        self.bowling.roll(5)
        self.assertEqual(46, self.bowling.score())

    def test_spare_n_strike(self):
        self.bowling.roll(6)
        self.bowling.roll(4)
        self.bowling.roll(10)
        self.bowling.roll(4)
        self.bowling.roll(1)
        self.assertEqual(40, self.bowling.score())

    def test_full_strike(self):
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.bowling.roll(10)
        self.assertEqual(300, self.bowling.score())


if __name__ == "__main__":
    unittest.main()
