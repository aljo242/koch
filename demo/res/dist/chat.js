"use strict";
const DEFAULT_NAME = "anon";
const DEFAULT_DECODING = "utf-8";
// TODO MAKE CheckHTTPS() func
const currentURL = window.location.href;
console.log(currentURL);
let websocketPrefix = "";
if (currentURL.includes("http:/")) {
    console.log("USING HTTP");
    websocketPrefix = "ws://";
}
if (currentURL.includes("https://")) {
    console.log("USING HTTPS");
    websocketPrefix = "wss://";
}
console.log(websocketPrefix);
if (!("TextEncoder" in window)) {
    alert("Sorry, this browser does not support TextEncoder!");
}
let encoder = new TextEncoder();
if (!("TextDecoder" in window)) {
    alert("Sorry, this browser does not support TextEncoder!");
}
let decoder = new TextDecoder(DEFAULT_DECODING);
function encode(msg) {
    return encoder.encode(msg).buffer;
}
function decode(buf) {
    return decoder.decode(buf);
}
if (!("WebSocket" in window)) {
    alert("Sorry, this browser does not support WebSockets!");
}
class User {
    constructor(name, conn) {
        console.log("Creating new User...");
        console.log(this);
        this.userName = name;
        this.conn = conn;
        this.signIn();
    }
    signIn() {
        const signInMessage = encode(`<b>${this.userName} signed in.</b>`);
        this.broadcast(signInMessage);
    }
    broadcast(buf) {
        this.conn.send(buf);
    }
    p2pSend(msg, target) {
        console.log(`Sending message to ${target}`);
        console.log(msg);
        this.conn.send(msg);
    }
}
let loginPopUpOpen = false;
function openPopUpForm() {
    let form = document.getElementById("popUpForm");
    if (form == null) {
        console.log("Unable to access login form!");
        return;
    }
    form.style.display = "block";
    loginPopUpOpen = true;
    console.log("opening login pop up");
}
function closePopUpForm() {
    let form = document.getElementById("popUpForm");
    if (form == null) {
        console.log("Unable to access login form!");
        return;
    }
    form.style.display = "none";
    loginPopUpOpen = false;
    console.log("closing login pop up");
}
//Since you know what type you are expecting, but the type 
//system can't, to get this to work, you have to tell Typescript 
//what type of element you expect to be selecting. You would do that 
//through casting the type of the selected element as follows: 
//const inputElement = <HTMLInputElement> document.getElementById("food-name-val"); 
//or 
//const inputElement = document.getElementById("food-name-val") as HTMLInputElement;
//window.onclick = (event : MouseEvent) => {
//    let modal = document.getElementById("popUpForm")!;
//    if (event.target == modal) {
//        closePopUpForm();
//    }
//}
function appendLog(item) {
    let log = document.getElementById("log");
    const doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    log.appendChild(item);
    if (doScroll) {
        log.scrollTop = log.scrollHeight - log.clientHeight;
    }
}
window.onload = () => {
    openPopUpForm();
    let conn;
    let user;
    let msg = document.getElementById("msg");
    if (window["WebSocket"]) {
        conn = new WebSocket(websocketPrefix + document.location.host + "/chat/ws");
        conn.binaryType = "arraybuffer";
        conn.onclose = () => {
            let item = document.createElement("div");
            item.innerHTML = "<b>Connection to server closed.</b>";
            appendLog(item);
            console.log("closing WS...");
        };
        conn.onmessage = (evt) => {
            //console.log(evt);
            let buf = evt.data;
            //let message : string[] = decode(evt.data).split("\n");
            let item = document.createElement("div");
            item.innerHTML = `${buf}`;
            appendLog(item);
        };
        conn.onclose = (evt) => {
            console.log(evt);
        };
    }
    else {
        let item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
        return;
    }
    let signInButton = document.getElementById("signInButton");
    let signIn = () => {
        let userName = document.getElementById("chatname");
        if (userName.value == "") {
            console.error("USER INPUT ERROR");
            userName.value = DEFAULT_NAME;
        }
        console.log(`user submitted to login form as: ${userName.value}`);
        user = new User(userName.value, conn);
        closePopUpForm();
    };
    signInButton.onclick = signIn;
    //let formKeyCallback = (ev: KeyboardEvent) => {
    //    if (loginPopUpOpen) {
    //        const enterCode = "Enter";
    //        const keyCode = ev.key;
    //        if (keyCode === enterCode) {
    //            // get form and submit it
    //            console.log(`User hit ${enterCode}`)
    //            signIn()
    //        }
    //    }
    //}
    //window.onkeypress = formKeyCallback;
    let msgForm = document.getElementById("send_msg_form");
    msgForm.onsubmit = () => {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        const messageWithName = encode(`${user.userName}: ${msg.value}`);
        user.broadcast(messageWithName);
        msg.value = "";
        return false;
    };
};
//# sourceMappingURL=chat.js.map