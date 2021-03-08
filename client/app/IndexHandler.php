<?php

namespace App;

use App\Util\Http;

class IndexHandler
{
    public function foo()
    {
        $baseWord = "AJAX is a developer's dream, because you can";
        $word = base64_encode($baseWord);
        $links = file_get_contents("data/links.json");
        $requestLink = sprintf(
            "http://localhost:8081/links-with-word/?key=%s&word=%s&links=%s",
            REQUEST_KEY,
            $word,
            implode(",", json_decode($links))
        );
        $linksWithWord = Http::get($requestLink);
        $linksWithWord = array_filter($linksWithWord);
        $linksWithWord = implode(", ", $linksWithWord);
        if (!$linksWithWord) {
            $linksWithWord = "none";
        }

        echo sprintf("Links with '%s': %s%s", $baseWord, $linksWithWord, PHP_EOL);
    }
}
