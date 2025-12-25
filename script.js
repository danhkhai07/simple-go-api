document.addEventListener("DOMContentLoaded", () => {
    const button = document.getElementById("magicButton");

    button.addEventListener("click", () => {
        alert("You clicked the magic button! ✨");
        // Add a cool background color change on click
        document.body.style.background = `linear-gradient(${Math.random()*360}deg, #${Math.floor(Math.random()*16777215).toString(16)}, #${Math.floor(Math.random()*16777215).toString(16)})`;
    });
});

