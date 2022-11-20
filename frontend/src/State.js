class UserStatus {
    static isEntered = false;

    static setStatus(status) {
        this.isEntered = status;
    }

    static getStatus() {
        return this.isEntered;
    }

    static switchStatus() {
        this.isEntered = !this.isEntered; 
    }
}

export default UserStatus