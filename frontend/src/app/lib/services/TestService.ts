class TestService {
    async test(): Promise<any> {
        const invite = await fetch("http://localhost:3202/invite/018e4c8f-21e4-73b7-8981-1ae97d90bfb0")
        return invite;
    }
}

export default new TestService();