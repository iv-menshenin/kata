import unittest

from factor import factor


class TestFactor(unittest.TestCase):

    def test_factor(self):
        self.assertEqual(factor(1), [1])
        self.assertEqual(factor(2), [2])
        self.assertEqual(factor(3), [3])
        self.assertEqual(factor(4), [2, 2])
        self.assertEqual(factor(26), [2, 13])


if __name__ == "__main__":
    unittest.main()
