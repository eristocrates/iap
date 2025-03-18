# Define the string to replace and the replacement string
$searchString = "/mnt/c/Users/eristocrates/notes/graphviz/"
$replaceString = "" # Replace with an empty string or your desired replacement

# Process each .dot file recursively
Get-ChildItem -Filter *.dot -Recurse | ForEach-Object {
    # Read the content of the .dot file
    $content = Get-Content -Path $_.FullName -Raw

    # Replace the search string with the replacement string
    $updatedContent = $content -replace [regex]::Escape($searchString), $replaceString

    # Write the updated content back to the file
    Set-Content -Path $_.FullName -Value $updatedContent

    # Convert the updated .dot file to .svg
    $outputFile = [System.IO.Path]::ChangeExtension($_.FullName, ".svg")
    dot -Tsvg $_.FullName -o $outputFile

    Write-Output "Processed $($_.FullName) and saved as $outputFile"
}