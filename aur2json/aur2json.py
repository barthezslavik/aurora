import json

dsl = """
Controller.UserProfile ->
    GetProfile(userId: Integer) -> { GetUser(userId), GetPostStats(userId) }
    UpdateAge(userId: Integer, newAge: Integer) -> ValidateAge(newAge) ? UpdateUserAge(userId, newAge) : Error("Invalid age")
    GetPostStats(userId: Integer) -> { Count(Post.findAllByUserId(userId)), avgLength(Post.findAllByUserId(userId)) }
    GetUser(userId: Integer) -> User.find(userId) or Error("Not found")
    ValidateAge(age: Integer) -> age in range(0, 150)
    UpdateUserAge(userId: Integer, age: Integer) -> User.find(userId).update(age: age)
    AvgLength(posts: Post[]) -> posts.empty() ? 0 : SumLengths(posts) / posts.count()
"""

def parse_dsl(dsl):
    lines = dsl.strip().split('\n')[1:]  # Skip the first line (Controller declaration)
    program = {
        "type": "Program",
        "body": [
            {
                "type": "StructDeclaration",
                "name": "UserProfileController",
                "fields": []
            },
            {
                "type": "StructDeclaration",
                "name": "User",
                "fields": [
                    {"name": "ID", "type": "int"},
                    {"name": "Age", "type": "int"}
                ]
            },
            {
                "type": "StructDeclaration",
                "name": "Post",
                "fields": [
                    {"name": "UserID", "type": "int"}
                ]
            }
        ]
    }

    for line in lines:
        parts = line.strip().split('->')
        method_def = parts[0].strip()
        method_name, params = method_def.split('(')
        method_name = method_name.strip()
        params = params.strip(') ')
        param_list = [{'name': p.split(':')[0].strip(), 'type': map_type(p.split(':')[1].strip())} for p in params.split(', ')]

        returns = infer_return_types(line)

        method_node = {
            'type': 'MethodDeclaration',
            'receiver': 'UserProfileController',
            'name': method_name,
            'parameters': param_list,
            'returns': returns
        }
        program['body'].append(method_node)

    return program

def map_type(param_type):
    return {
        "Integer": "int",
        "Post[]": "[]Post"
    }.get(param_type, param_type)

def infer_return_types(line):
    if "Error" in line:
        return ["error"]
    if "GetProfile" in line or "GetUser" in line:
        return ["*User", "error"]
    if "GetPostStats" in line:
        return ["PostStats", "error"]
    if "ValidateAge" in line:
        return ["bool"]
    if "AvgLength" in line:
        return ["float64"]
    return []

ast_json = parse_dsl(dsl)
print(json.dumps(ast_json, indent=4))

# Save to file
with open('../json2go/generated_ast.json', 'w') as f:
	json.dump(ast_json, f, indent=4)