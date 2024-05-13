using Azure.Identity;
using System.Net.Http.Json;
var operation = args[0];
if (operation == "get")
{
    //using var file = File.CreateText("log.txt");
    //file.WriteLine($"{DateTime.Now}:OperationName:{operation}");
    var registery = Console.ReadLine();
    file.WriteLine($"{DateTime.Now}:RegisteryName:{registery}");
    var credential = new DefaultAzureCredential(new DefaultAzureCredentialOptions() { ExcludeManagedIdentityCredential = true });
    var tokenResult = await credential.GetTokenAsync(new Azure.Core.TokenRequestContext(new[] { "https://containerregistry.azure.net/.default" }), CancellationToken.None);
    file.WriteLine($"{DateTime.Now}:ARMToken:{tokenResult.Token}");
    HttpClient client = new HttpClient();
    Dictionary<string, string> payload = new Dictionary<string, string>
    {
        { "grant_type", "access_token" },
        { "service", registery },
        { "access_token", tokenResult.Token }
    };
    var content = new FormUrlEncodedContent(payload);
    var response = client.PostAsync($"https://{registery}/oauth2/exchange", content);
    var tokenResponse = await response.Result.Content.ReadFromJsonAsync<TokenResponse>();
    var output = $"{{\"Username\": \"<token>\",\"Secret\": \"{tokenResponse.refresh_token}\"}}";
    file.WriteLine($"{DateTime.Now}:output:{output}");
    Console.WriteLine(output);
}
public class TokenResponse
{
    public string refresh_token { get; set; }
}