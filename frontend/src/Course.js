class Course {
    static id = 0;
    name = "";
    img = "";
    ref = "";
    cost = 0;

    constructor(name, img, ref, cost) {
        this.name = name;
        this.img = img;
        this.ref = ref;
        this.cost = cost;
    }
}

export default Course;