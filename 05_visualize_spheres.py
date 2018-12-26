import sys
import numpy as np

from numpy.linalg import norm

from hitable_list import HitableList
from ray import Ray
from sphere import Sphere


RESULT_FILE = "resources/chap05-2.ppm"


def color(r, world):
    record = {'t': 0, 'p': np.zeros(3), 'normal': np.zeros(3)}
    hit = world.hit(r, 0.0, 999999, record)
    if hit[0]:
        return 0.5 * (hit[1]['normal'] + 1)
    else:
        unit_direction = r.direction / norm(r.direction)
        t = 0.5 * (unit_direction[1] + 1.0)
        return (1.0 - t) * np.array([1.0, 1.0, 1.0]) + t * np.array([0.5, 0.7, 1.0])


def lerp(nx, ny):
    result = ["P3", f"{nx} {ny}", "255"]

    lower_left_corner = np.array([-2.0, -1.0, -1.0])
    horizontal = np.array([4.0, 0.0, 0.0])
    vertical = np.array([0.0, 2.0, 0.0])
    origin = np.zeros(3)

    s1 = Sphere(np.array([0, 0, -1]), 0.5)
    s2 = Sphere(np.array([0, -100.5, -1]), 100)
    h = HitableList([s1, s2])

    matrix = list()
    for j in range(ny - 1, -1, -1):
        for i in range(nx):
            u = i / nx
            v = j / ny
            r = Ray(origin, lower_left_corner + u * horizontal + v * vertical)
            c = color(r, h)
            ir = int(255.99 * c[0])
            ig = int(255.99 * c[1])
            ib = int(255.99 * c[2])
            matrix.append([ir, ig, ib])
            result.append(f"{ir} {ig} {ib}")

    p = np.array(matrix)
    print(p.shape)
    return result


def main():
    with open(RESULT_FILE, mode='a', encoding='utf8', newline='\n') as f:
        for line in lerp(200, 100):
            f.write(line + "\n")


if __name__ == '__main__':
    main()