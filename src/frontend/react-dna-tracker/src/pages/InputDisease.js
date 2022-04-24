export default function InputDisease() {
    function handleSubmit(e) {
        e.preventDefault();
    }
    return (
        <section className="container">
            <h1 className="header">Input Disease</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Disease name:
                    </label>
                    <input type="text" placeholder="Input the disease name" />
                </div>
                <div className="form-control">
                    <label>
                        DNA sequence:
                    </label>
                    <input type="file" />
                </div>
                <input type='submit' value='Submit Disease' className = 'btn'/>
            </form>
        </section>
    )
}