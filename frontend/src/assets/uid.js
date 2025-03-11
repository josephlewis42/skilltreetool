var uid = 0;

export function uniqueId() {
    uid++;
    return "id-" + uid;
}

