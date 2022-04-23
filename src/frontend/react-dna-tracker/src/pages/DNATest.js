export default function DNATest() {
    return (
        <section>
            <h1>DNA Test</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Patient name:
                        <input type="text" placeholder="Input the patient name" />
                    </label>
                </div>
                <div className="form-control">
                    <label>
                        DNA Sequence:
                        <input type="file"/>
                    </label>
                </div>
                <div>
                    <select>
                        <option>Penyakit 1</option>
                        <option>Penyakit 2</option>
                        <option>Penyakit 3</option>
                        <option>Penyakit 4</option>
                    </select>
                </div>

                <input type='submit' value='Submit Test' className = 'btn btn-block'/>
            </form>
        </section>
    )
}

function handleSubmit(e) {
    e.preventDefault();
}