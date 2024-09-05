function main() {
	const buttonSubmit = document.getElementById("button-submit")

	buttonSubmit.addEventListener("click", async () => {
		try {
			const link = document.getElementById("input-link");
			if (!link.value) {
				throw new Error("You forgot the link, bro 💀")
			}

			const status = document.getElementById("status")
			status.textContent = "in progress..."
			status.style.backgroundColor = "transparent";

			const url = await fetch("/api/download", {
				method: "POST",
				body: JSON.stringify({
					"link": link.value,
				})
			});
			const obj = await url.json();
			if (!obj.success) {
				throw new Error(obj.message)
			}

			const download = document.createElement("a");
			download.href = obj.url;
			download.download = obj.filename;
			document.body.appendChild(download);
			download.click();
			document.body.removeChild(download);

			status.textContent = "";
		} catch (error) {
			const status = document.getElementById("status");
			status.textContent = error.message;
			status.style.backgroundColor = "rgba(255, 0, 0, 0.5)"
			status.style.paddingLeft = "5px"
			status.style.paddingRight = "5px"
			return;
		}
	})
}
main();
