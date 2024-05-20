// Generate a valid token for testing purposes.
export function generateToken(): string {
    const payload = {
        sub: '1234567890',
        name: 'John Doe',
        exp: new Date().getTime() / 1000 + 3600,
    };

    const base64Url = Buffer.from(JSON.stringify(payload)).toString('base64');
    return `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.${base64Url}.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`;
}