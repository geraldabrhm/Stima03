export default function InputDisease() {
    function handleSubmit(e) {
        e.preventDefault();
    }
    return (
        <section>
            <h1>Input Disease</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Disease name:
                        <input type="text" placeholder="Input the disease name" />
                    </label>
                </div>
                <div className="form-control">
                    <label>
                        DNA Sequence
                        <input type="file" name="name" />
                    </label>
                </div>
                <input type='submit' value='Submit Disease' className = 'btn btn-block'/>
            </form>
        </section>
    )
}