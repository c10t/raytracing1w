import numpy as np


class Ray:
    origin = np.zeros(3)
    direction = np.zeros(3)

    def __init__(self, origin, direction):
        self.origin = origin
        self.direction = direction

    def point_at(self, t):
        return self.origin + t * self.direction


if __name__ == '__main__':
    o = np.array([0, 1, 2])
    d = np.array([3, 4, 5])
    r = Ray(o, d)
    print(f"{r.origin} + 3 * {r.direction} = {r.point_at(3)}")
