str1 = [
    {
    id: "47",
    msg: "hello",
    },
    {
    id: "47",
    msg: "hello",
    },
]
str2 = {
    id: "47",
    msg: "hello",
}

// console.log(str1[0])
// console.log(str2)

for (err of str1) {
    if (err.id == str2.id && err.msg == str2.msg) {
        console.log("Pass 1")
    }
}
// if (str1[0].msg === str2.msg) {
//     console.log("Pass 2")
// }
if (str1.includes(str2)) {
    console.log("Pass 3")
}
if (str2 in str1) {
    console.log("Pass 4")
}