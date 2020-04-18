package py

const PythonServerTemplate = `
from flask import render_template
import connexion

# Create the application instance
app = connexion.App(__name__, specification_dir='spec')

# Read apiSpec file to setup api endpoints
app.add_api('spec.yml')

@app.route('/home')
def home():
    """
    This function just responds to the browser ULR
    :return:        the rendered template 'index.html'
    """
    return render_template('index.html')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port={{.}}, debug=True)`
