function main() {
	const buttonSubmit = document.getElementById("button-submit")

	buttonSubmit.addEventListener("click", async () => {
		const link = document.getElementById("input-link").value;

		const response = await fetch("/api/download", {
			method: "POST",
			body: {
				link: link,
			}
		})
		console.log(await response.json());
	})
}
main();
