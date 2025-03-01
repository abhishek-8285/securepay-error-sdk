import json

with open("src/errors/errorCodes.json", "r") as file:
    error_codes = json.load(file)

class ErrorHandler:
    @staticmethod
    def get_error(error_key):
        return error_codes.get(error_key, error_codes["SYSTEM_INTERNAL_ERROR"])

    @staticmethod
    def error_response(error_key):
        error = ErrorHandler.get_error(error_key)
        return {
            "status": "FAILED",
            "error": {
                "code": error["code"],
                "message": error["message"]
            }
        }
