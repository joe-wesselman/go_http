document.addEventListener("DOMContentLoaded", () =>{
    const btn = document.getElementById("clickMe");
    const output = document.getElementById("output");

    btn.addEventListener("click", () =>{
        const time = new Date().toLocaleTimeString();
        output.textContent = `Button clicked at ${time}`;
        btn.style.background = "#00c853";
        output.classList.add("visible");

        setTimeout(() =>{
            btn.style.background = "#008cff";
            output.classList.remove("visible");
            setTimeout(() =>{
                output.textContent = "";
            }, 1000);
        }, 2000);
    });
});