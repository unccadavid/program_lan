import unittest
import tf_05
import copy

class TestTF05(unittest.TestCase):
    def test_read_file_non_idempotent(self):
        path = "text_file.txt"
        tf_05.read_file(path)
        answer1 = copy.deepcopy(tf_05.data)
        tf_05.read_file(path)
        answer2 = copy.deepcopy(tf_05.data)
        self.assertNotEqual(answer1, answer2)
        #因為readfile每次執行，
        #會添加當前開啟的檔案內容到data，
        #所以只要重複call了這個，data就會不斷新增額外的內容
        #因此為non-idempotent

    def test_scan_non_idempotent(self):
        tf_05.data = ["this","is","a","test"]
        tf_05.scan()
        answer1 = copy.deepcopy(tf_05.words)
        tf_05.scan()
        answer2 = copy.deepcopy(tf_05.words)
        self.assertNotEqual(answer1, answer2)
        #因為scan每次執行，
        #會將當前list的內容轉為string並添加到words，
        #所以只要重複call了這個，word就會不斷新增string進去
        #因此為non-idempotent

    def test_frequencies_non_idempotent(self):
        tf_05.words=["program","program","program","program","test","test"]
        tf_05.frequencies()
        answer1 = copy.deepcopy(tf_05.word_freqs)
        tf_05.frequencies()
        answer2 = copy.deepcopy(tf_05.word_freqs.copy)
        self.assertNotEqual(answer1, answer2)
        #因為frequencies每次執行，
        #會將word_freq添加上，這次word中出現的字，將其計數+1
        #所以只要重複call了這個，word_freqs中的數字就會不斷上升
        #因此為non-idempotent
        
if __name__ == '__main__':
    unittest.main()