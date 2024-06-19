from textx import metamodel_from_file
import json, argparse

def model_to_dict(model):
    return {
        "namespace": model.namespace,
        "relations": [{
           "name": relation.relationName,
           "additional": [{ 
                "type": t.__class__.__name__,
                "children": [{
                    "child": "this" if child.this is not None else child.existingRelation.relationName
                } for child in t.children]
           } for t in relation.additionalInfo.type] if relation.additionalInfo is not None else []
        } for relation in model.relations]
    }

def read_model(path):
    ret = ""
    with open(path, "r") as f:
        ret = f.read()
    return ret

def export_dict(model_dict):
    with open("template.json", "w") as f:
        f.write(json.dumps(model_dict))

if __name__=="__main__":
    parser = argparse.ArgumentParser(description='Parse a DSL model using a specified grammar.')
    parser.add_argument('--grammar', type=str, required=True, help='Path to the grammar file')
    parser.add_argument('--model', type=str, required=True, help='Path to the model file')
    args = parser.parse_args()

    model_str = read_model(args.model)
    mm = metamodel_from_file(args.grammar)
    model = mm.model_from_str(model_str)
    model_dict = model_to_dict(model)

    #print(json.dumps(model_dict))

    export_dict(model_dict)