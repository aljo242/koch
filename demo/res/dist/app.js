"use strict";
const black = "#000000";
class ColorScheme {
    constructor() {
        this.article_background = black;
        this.article_opacity = "1.0";
        this.nav_background = black;
        this.nav_header = black;
        this.nav_opacity = "1.0";
        this.body_background = black;
        this.hr = black;
        this.a_nav_link = black;
        this.a_nav_visited = black;
        this.a_nav_hover = black;
        this.a_nav_active = black;
        this.a_text_link = black;
        this.a_text_visited = black;
        this.a_text_hover = black;
        this.a_text_active = black;
        this.main_text = black;
    }
}
let scheme1 = new ColorScheme;
scheme1.article_background = "#f4a261";
scheme1.nav_background = "#264653";
scheme1.nav_header = "#e76f51";
scheme1.body_background = black;
scheme1.hr = black;
scheme1.a_nav_link = "#14ff14";
scheme1.a_text_link = "#e9c46a";
let scheme2 = new ColorScheme;
scheme2.article_background = "#f0e68c";
scheme2.nav_background = "#adff2f";
scheme2.nav_header = "#ffc0cb";
scheme2.body_background = black;
scheme2.hr = black;
scheme2.a_nav_link = "#f700ff";
scheme2.a_text_link = "#f700ff";
let scheme3 = new ColorScheme;
scheme3.article_background = "#e5e5e5";
scheme3.nav_background = "#e5e5e5";
scheme3.nav_header = "#fb8500";
scheme3.body_background = black;
scheme3.hr = black;
scheme3.a_nav_link = "#52b788";
scheme3.a_text_link = "#52b788";
let scheme4 = new ColorScheme;
scheme4.article_background = "#e2fdff";
scheme4.nav_background = "#bfd7ff";
scheme4.nav_header = "#5465ff";
scheme4.body_background = black;
scheme4.hr = black;
scheme4.a_nav_link = "#c32f27";
scheme4.a_text_link = "#c32f27";
let scheme5 = new ColorScheme;
scheme5.article_background = "#f28482";
scheme5.article_opacity = ".975";
scheme5.nav_background = "#84a59d";
scheme5.nav_header = "#f7ede2";
scheme5.nav_opacity = ".975";
scheme5.body_background = black;
scheme5.hr = black;
scheme5.a_nav_link = "#14ff14";
scheme5.a_text_link = "#14ff14";
let scheme6 = new ColorScheme;
scheme6.article_background = "#ddc9b4";
scheme6.article_opacity = ".95";
scheme6.nav_background = "#bcac9b";
scheme6.nav_header = "#c17c74";
scheme6.nav_opacity = ".90";
scheme6.body_background = black;
scheme6.hr = black;
scheme6.a_nav_link = "#df2935";
scheme6.a_text_link = "#df2935";
let scheme7 = new ColorScheme;
scheme7.article_background = "#dddf00";
scheme7.article_opacity = ".95";
scheme7.nav_background = "#2b9348";
scheme7.nav_header = "#007f5f";
scheme7.nav_opacity = ".90";
scheme7.body_background = black;
scheme7.hr = black;
scheme7.a_nav_link = "#001d3d";
scheme7.a_text_link = "#001d3d";
// Register Service Worker
if ("serviceWorker" in navigator) {
    window.addEventListener("load", function () {
        navigator.serviceWorker.register("/serviceWorker.js").then(function (registration) {
            // Registration was successful
            console.log("ServiceWorker registration successful with scope: ", registration.scope);
        }, function (err) {
            // registration failed :(
            console.log("ServiceWorker registration failed: ", err);
        });
    });
}
const msg = "Hello web";
function printMessage(msg) {
    console.log(msg);
}
function applyColorScheme(scheme) {
    //console.log(scheme)
    let header = document.getElementById("header");
    header.style.backgroundColor = scheme.nav_header;
    let nav = document.getElementById("nav");
    nav.style.backgroundColor = scheme.nav_background;
    nav.style.opacity = scheme.nav_opacity;
    let mainArticleBody = document.getElementById("mainArticleBody");
    mainArticleBody.style.backgroundColor = scheme.article_background;
    mainArticleBody.style.opacity = scheme.article_opacity;
    //let navLinks = document.getElementsByClassName("navLink")!;
    //for (let item  of navLinks) {
    //  let hmtlItem = <HTMLElement>item;
    //  hmtlItem.style.color = scheme.a_nav_link;
    //}
    //let textLinks = document.getElementsByClassName("textLink")!;
    //for (let item  of textLinks) {
    //  let hmtlItem = <HTMLElement>item;
    //  hmtlItem.style.color = scheme.a_text_link;
    //}
    let links = document.getElementsByTagName("a");
    for (let item of links) {
        item.style.color = scheme.a_nav_link;
    }
}
const colorSchemes = [scheme1, scheme2, scheme3, scheme4, scheme5, scheme6, scheme7];
let currentSelection = 0;
let prevSelection = 0;
function applyRandomColorScheme() {
    let random = Math.floor(Math.random() * colorSchemes.length);
    // make sure we always choose a unique one
    while (random === currentSelection || random === prevSelection) {
        random = Math.floor(Math.random() * colorSchemes.length);
    }
    prevSelection = currentSelection;
    currentSelection = random;
    console.log(currentSelection);
    applyColorScheme(colorSchemes[currentSelection]);
}
applyColorScheme(colorSchemes[currentSelection]);
const homeButton = document.getElementById("secretButton");
homeButton.onclick = () => {
    applyRandomColorScheme();
};
//# sourceMappingURL=app.js.map