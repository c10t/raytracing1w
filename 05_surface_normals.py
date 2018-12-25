import math
import numpy as np
from numpy.linalg import norm

from ray import Ray

RESULT_FILE = "resources/chap05.ppm"


def hit_sphere(center, radius, r):
    oc = r.origin - center
    a = np.dot(r.direction, r.direction)
    b = 2.0 * np.dot(oc, r.direction)
    c = np.dot(oc, oc) - radius * radius
    discriminant = b * b - 4 * a * c
    if discriminant < 0:
        return -1.0
    else:
        return (-b - math.sqrt(discriminant)) / (2.0 * a)


def color(r):
    t = hit_sphere(np.array([0, 0, -1]), 0.5, r)
    if t > 0.0:
        pc = r.point_at(t) - np.array([0, 0, -1])
        N =  pc / norm(pc)
        return 0.5 * (N + 1)

    unit_direction = r.direction / norm(r.direction)
    t = 0.5 * (unit_direction[1] + 1.0)
    return (1.0 - t) * np.array([1.0, 1.0, 1.0]) + t * np.array([0.5, 0.7, 1.0])


def lerp(nx, ny):
    result = ["P3", f"{nx} {ny}", "255"]

    lower_left_corner = np.array([-2.0, -1.0, -1.0])
    horizontal = np.array([4.0, 0.0, 0.0])
    vertical = np.array([0.0, 2.0, 0.0])
    origin = np.zeros(3)

    matrix = list()
    for j in range(ny - 1, -1, -1):
        for i in range(nx):
            u = i / nx
            v = j / ny
            r = Ray(origin, lower_left_corner + u * horizontal + v * vertical)
            c = color(r)
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
