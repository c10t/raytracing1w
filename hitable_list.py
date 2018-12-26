class HitableList:
    items = []
    size = 0

    def __init__(self, hitables):
        self.items = hitables
        self.size = len(hitables)

    def hit(self, r, t_min, t_max, record):
        new_record = record.copy()
        hit_anything = False
        closest_so_far = t_max
        for item in self.items:
            h = item.hit(r, t_min, closest_so_far, new_record)
            if h[0]:
                hit_anything = True
                closest_so_far = h[1]['t']
                new_record = h[1].copy()

        return hit_anything, new_record
