const input = document.querySelector("input")
const button = document.querySelector("button")

window.onload = () => {
  input.value = ""
  input.focus()
}

async function shorten() {
    let response = await fetch("/shorten", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ url: input.value })
    })

    let data = await response.json()

    input.value = window.location.origin + "/" + data.id
    button.innerText = "Copy"

    navigator.clipboard.writeText(input.value)
}

window.addEventListener("keydown", (event) => {
    if (event.key === "Enter" && input.value && input == document.activeElement) {
        shorten()
    }
})

button.addEventListener("click", async event => {
  if (button.innerText == "Copy") {
    navigator.clipboard.writeText(input.value)
    button.innerText = "Copied!"

    setTimeout(() => {
      button.innerText = "Visit"
    }, 1000)
  } else if (button.innerText == "Shorten" && input.value) {
    await shorten(event)
  } else if (button.innerText == "Visit") {
    window.open(input.value, "_blank")
    input.value = ""
    button.innerText = "Shorten"
  }
})

input.addEventListener("input", event => {
  if (button.innerText == "Copy") {
    input.value = ""
    button.innerText = "Shorten"
  }
})
