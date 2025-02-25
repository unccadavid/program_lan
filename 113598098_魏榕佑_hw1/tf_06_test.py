import unittest
from tf_06 import *

class TestTF06(unittest.TestCase):
    def test_filter_chars_and_normalize(self):
        input = "This is a test. This test is only a test."
        expected = "this is a test this test is only a test "

        result = filter_chars_and_normalize(input)

        self.assertEqual(expected, result)

    def test_scan(self):
        input = "this is a test this test is only a test "
        expected = ["this","is","a","test","this","test","is","only","a","test"]
        result = scan(input)
        self.assertEqual(expected, result)

    def test_remove_stop_words(self):
        input = ["this","is","a","test","this","test","is","only","a","test"]
        path = "stop_words.txt"
        expected = ["test","test","test"]
        stop_word_filter = remove_stop_words(input)
        result = stop_word_filter(path)
        self.assertEqual(expected, result)

    def test_frequencies(self):
        input = ["test","test","test","program","program","program","program"]
        expected = {"test":3,"program":4}
        result = frequencies(input)
        self.assertEqual(expected, result)

    def test_sort(self):
        input = {"test":3,"program":4}
        expected = [("program",4),("test",3)]
        result = sort(input)
        self.assertEqual(expected, result)

if __name__ == "__main__":
    unittest.main()