import unittest

from wrapper import wrapper


class TestWrapper(unittest.TestCase):

    def test_wrapper(self):
        self.assertEqual([], wrapper("", 5))
        self.assertEqual(["fits"], wrapper("fits", 4))
        self.assertEqual(["fi", "ts"], wrapper("fits", 2))
        self.assertEqual(["its", "not", "fit"], wrapper("its not fit", 3))
        self.assertEqual(["its", "not", "fit"], wrapper("its not fit", 4))
        self.assertEqual(["its", "not", "fit"], wrapper("its not fit", 5))
        self.assertEqual(["its not", "fit"], wrapper("its not fit", 9))


if __name__ == "__main__":
    unittest.main()
