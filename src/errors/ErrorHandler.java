import com.fasterxml.jackson.databind.ObjectMapper;
import java.io.File;
import java.io.IOException;
import java.util.Map;

public class ErrorHandler {
    private static Map<String, Map<String, Object>> errorCodes;

    static {
        try {
            ObjectMapper objectMapper = new ObjectMapper();
            errorCodes = objectMapper.readValue(new File("src/errors/errorCodes.json"), Map.class);
        } catch (IOException e) {
            throw new RuntimeException("Failed to load error mappings", e);
        }
    }

    public static Map<String, Object> getError(String errorKey) {
        return errorCodes.getOrDefault(errorKey, errorCodes.get("SYSTEM_INTERNAL_ERROR"));
    }
}
