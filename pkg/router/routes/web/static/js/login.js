async function login() {
    const username = document.getElementById('username').value.trim()
    const password = document.getElementById('password').value.trim()
    if (!username || !password) {
        alert('Please enter a username and password')
        return
    }
    try {
        const url =
            '/api/login?username=' +
            encodeURIComponent(username) +
            '&password=' +
            encodeURIComponent(password)
        const response = await fetch(url)
        if (response.status != 200) {
            const text = await response.text()
            alert(text)
            return
        }
        window.location.href = '/'
    } catch (error) {
        console.error(error)
    }
}
