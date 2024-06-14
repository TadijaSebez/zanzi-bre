from textx import metamodel_from_str, get_children_of_type
import json


grammar = """
Model:
    'ime' ':' namespace=STRING    
    relations*=Relation
;

Relation:
    'odnos' '{'
        'ime' ':' relationName=STRING
        (additionalInfo=AdditionalInfo)?
    '}'
;

AdditionalInfo:
    'dodatak' '{'
        type+=Type
    '}'
;

Union:
    'bilo_koji' '{'
        children+=Child
    '}'
;

Intersection:
    'svaki' '{'
        children+=Child
    '}'
;

Exclusion:
    'jedan_od' '{'
        children+=Child
    '}'
;

Child:
    'dete' '{'
        ( this=This
        | existingRelation=ExistingRelation)
    '}'
;

This:
    'ovaj' '{}'
;

ExistingRelation:
    'postojeci_odnos' '{' 
        'odnos' ':' relationName=STRING
    '}'
;

Type:
    Union | Intersection | Exclusion
;
"""

model_str = """
ime: "doc"

odnos { 
    ime: "owner" 
}

odnos {
    ime: "editor"
    dodatak {
        bilo_koji {
            dete { ovaj {} }
            dete { postojeci_odnos { odnos: "owner" } }
        }
    }
}

odnos {
    ime: "viewer"
    dodatak {
        svaki {
            dete { ovaj {} }
            dete { postojeci_odnos { odnos: "editor" } }
        }
    }
}
"""

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
           } for t in relation.additionalInfo.type] if relation.additionalInfo is not None else {}
        } for relation in model.relations]
    }

if __name__=="__main__":
    mm = metamodel_from_str(grammar)
    model = mm.model_from_str(model_str)

    def cname(o):
        return o.__class__.__name__

    for relation in model.relations:
        print(relation.relationName)

    model_dict = model_to_dict(model)
    print(json.dumps(model_dict))