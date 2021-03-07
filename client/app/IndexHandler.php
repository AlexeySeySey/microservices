<?php

namespace App;

use App\Util\Http;

class IndexHandler
{
    public function foo()
    {
        $word = "foo";
        $links = include "../data/links.json";
        $requestLink = sprintf(
            "http://localhost:8081/links-with-word/?key=%s&word=&links=%s",
            REQUEST_KEY,
            $word,
            http_build_query($links)
        );

        $linksWithWord = Http::get($requestLink);

        // TODO: can't run PHP

        return $linksWithWord;

//        $linksWithWord = count($linksWithWord) > 0 ?
//            implode(", ", $linksWithWord) :
//            "none";
//
//        return sprintf("Links with '%s': %s", $word, $linksWithWord);
    }
}
