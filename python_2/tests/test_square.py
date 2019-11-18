from func import square

def test_squaring():
    assert square(2) == 4

def test_square_fail():
    assert square(4) == 15