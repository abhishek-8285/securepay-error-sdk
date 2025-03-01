const fs = require('fs');
const path = require('path');

// Load error mappings
const errorCodes = JSON.parse(fs.readFileSync(path.join(__dirname, 'errorCodes.json'), 'utf-8'));

class ErrorHandler {
    static getError(errorKey) {
        return errorCodes[errorKey] || errorCodes["SYSTEM_INTERNAL_ERROR"];
    }

    static errorResponse(res, errorKey) {
        const error = this.getError(errorKey);
        return res.status(error.httpStatus).json({
            status: "FAILED",
            error: {
                code: error.code,
                message: error.message
            }
        });
    }
}

const ErrorKeys = Object.keys(errorCodes).reduce((acc, key) => {
    acc[key] = key;
    return acc;
}, {});

module.exports = { ErrorHandler, ErrorKeys };
