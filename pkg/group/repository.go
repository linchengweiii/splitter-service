package group

type InMemoryRepository struct {
    groups []Group
}

func NewInMemoryRepository() *InMemoryRepository {
    return &InMemoryRepository{
        make([]Group, 0),
    }
}

func (r *InMemoryRepository) Create(group Group) error {
    r.groups = append(r.groups, group)
    return nil
}

func (r *InMemoryRepository) Read(id string) (Group, error) {
    for _, group := range r.groups {
        if group.Id == id {
            return group, nil
        }
    }
    return Group{}, nil
}

func (r *InMemoryRepository) Update(group Group) error {
    for i, g := range r.groups {
        if g.Id == group.Id {
            r.groups[i] = group
            return nil
        }
    }
    return nil
}

func (r *InMemoryRepository) Delete(id string) error {
    for i, group := range r.groups {
        if group.Id == id {
            r.groups = append(r.groups[:i], r.groups[i+1:]...)
            return nil
        }
    }
    return nil
}
