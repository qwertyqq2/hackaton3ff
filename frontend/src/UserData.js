class UserData {
    static id = 0;
    name;
    wallets = {};
    access;

    constructor(name, wallets, access) {
        this.name = name;
        this.wallets = wallets;
        this.access = access;
    }
}

export default UserData;