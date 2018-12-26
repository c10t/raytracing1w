import math
import numpy as np

from ray import Ray

class Sphere:
    center = np.zeros(3)
    radius = 0.0


    def __init__(self, center, radius):
        self.center = center
        self.radius = radius


    def hit(self, r, t_min, t_max, record):
        """
        check if the ray hit to the sphere or not.
        :param r: Ray
        :param t_min:
        :param t_max:
        :param record: { 't': scalar, 'p': vector, 'normal': vector }
        :return: (bool, record)
        """
        oc = r.origin - self.center
        a = np.dot(r.direction, r.direction)
        b = np.dot(oc, r.direction)
        c = np.dot(oc, oc) - self.radius * self.radius
        discriminant = b * b - a * c

        if discriminant > 0:
            # TODO: change variable name to cooler one
            temp1 = (-b - math.sqrt(discriminant)) / a
            if t_min < temp1 < t_max:
                new_record = dict()
                new_record['t'] = temp1
                new_record['p'] = r.point_at(temp1)
                new_record['normal'] = (new_record['p'] - self.center) / self.radius

                # original code in the book like...
                # record.t = temp
                # record.p = r.point_at(record.t)
                # record.normal = (record.p - self.center) / self.radius
                # return True
                return True, new_record
            # TODO: change variable name to cooler one
            temp2 = (-b + math.sqrt(discriminant)) / a
            if t_min < temp2 < t_max:
                new_record = dict()
                new_record['t'] = temp2
                new_record['p'] = r.point_at(temp2)
                new_record['normal'] = (new_record['p'] - self.center) / self.radius
                return True, new_record

        return False, record


if __name__ == '__main__':
    print("Test this class Sphere...")
    s0 = Sphere(np.array([0.0, 0.0, 0.0]), 3)
    r1 = Ray(np.array([0, 1, 2]), np.array([3, 4, 5]))
    h1 = s0.hit(r1, 0, 5, {})
    print(f"hit? {h1[0]}")
    print(f"record: {h1[1]}")
