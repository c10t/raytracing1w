import numpy as np

RESULT_FILE = "resources/chap02.ppm"


def hello_ppm(nx, ny):
    result = ["P3", f"{nx} {ny}", "255"]

    matrix = list()
    for j in range(ny - 1, -1, -1):
        for i in range(nx):
            r = i / nx
            g = j / ny
            b = 0.2
            ir = int(255.99 * r)
            ig = int(255.99 * g)
            ib = int(255.99 * b)
            matrix.append([ir, ig, ib])
            result.append(f"{ir} {ig} {ib}")

    p = np.array(matrix)
    print(p.shape)
    return result


def main():
    """
    with open(RESULT_FILE, mode='a', encoding='utf8', newline='\n') as f:
        for line in hello_ppm(200, 100):
            f.write(line + "\n")
    """
    hello_ppm(200, 100)


if __name__ == '__main__':
    main()
